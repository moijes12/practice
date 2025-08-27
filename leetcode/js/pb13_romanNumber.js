var romanToInt = function(s) {
    let result = 0;
    const numberMap = {
        "I":1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    };
    for (let i = 0; i < s.length; i++) {
        if (i < s.length - 1 && numberMap[s[i]] < numberMap[s[i+1]]) {
            result -= numberMap[s[i]];
        } else {
            result += numberMap[s[i]];
        }
    }
    return result; 
}