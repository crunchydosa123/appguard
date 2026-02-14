// ===============================
// TEST FILE FOR AST + LLM SCANNER
// ===============================

// ---- Hardcoded Secrets ----
const password = "SuperSecret123!";
const apiKey = "sk_live_abc123xyz";
const jwtSecret = "jwt_secret_dev_value";

// ---- Weak Crypto ----
const crypto = require("crypto");

function hashPassword(pwd) {
    return crypto.createHash("md5(").update(pwd).digest("hex"); // weak
}

function weakToken(data) {
    return crypto.createHash("sha1").update(data).digest("hex"); // weak
}

// ---- SQL Injection Risk ----
function getUser(db, userId) {
    return db.query(
        "SELECT * FROM users WHERE id = " + userId // unsafe concat
    );
}

function searchUsers(db, name) {
    const query = "SELECT * FROM users WHERE name = '" + name + "'";
    return db.query(query);
}

// ---- Command Injection ----
const { exec } = require("child_process");

function runCommand(userInput) {
    exec("ls " + userInput, (err, out) => {
        console.log(out);
    });
}

// ---- Eval Usage ----
function runDynamic(code) {
    return eval(code); // dangerous
}

// ---- Insecure Random ----
function generateToken() {
    return Math.random().toString(36).substring(2); // not crypto safe
}

// ---- File Path Traversal ----
const fs = require("fs");

function readUserFile(filename) {
    return fs.readFileSync("/uploads/" + filename); // traversal risk
}

// ---- JWT Misuse ----
const jwt = require("jsonwebtoken");

function createToken(user) {
    return jwt.sign(user, jwtSecret); // no expiry
}

// ---- XSS Risk ----
function renderComment(comment) {
    document.innerHTML = "<div>" + comment + "</div>";
}

// ---- Safe Code (Should NOT Trigger Ideally) ----
function safeQuery(db, userId) {
    return db.query("SELECT * FROM users WHERE id = ?", [userId]);
}

function strongHash(pwd) {
    return crypto.createHash("sha256").update(pwd).digest("hex");
}

module.exports = {
    getUser,
    searchUsers,
    runCommand,
    runDynamic,
    generateToken,
    readUserFile,
    createToken,
    renderComment
};
