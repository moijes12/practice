var columnNumber = function(columnTitle) {
    let result = 0;
    for (let c of columnTitle) {
        result = result * 26 + (c.charCodeAt(0) - 'A'.charCodeAt(0) + 1);
    }
    return result;
};