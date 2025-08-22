var isPalindrome = function(s) {
    s = s.replace(/[^a-zA-Z0-9]/g,'');
    s = s.toLowerCase();
    // Base case
    if (s.length === 0 || s.length === 1) return true;

    let startIndex = 0;
    let endIndex = s.length - 1;
    while (startIndex < endIndex) {
        if (s[startIndex] !== s[endIndex]) return false;
        startIndex++;
        endIndex--;
    }
    return true;
};