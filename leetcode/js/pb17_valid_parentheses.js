var isValid = function(s) {
    let brackStack = [];
    let bracketMap = {
        '}' : '{',
        ']': '[',
        ')' : '('
    };
    for (let c in s) {
        if (!bracketMap[c]) {
            brackStack.push(c);
        }
        else if (bracketMap[c] !== brackStack.pop()) {
            return false;
        }
    }
    return brackStack.length === 0;
};