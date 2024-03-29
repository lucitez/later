export const timeSince = (timestamp) => {
    var msPerSecond = 1000;
    var msPerMinute = 60 * 1000;
    var msPerHour = msPerMinute * 60;
    var msPerDay = msPerHour * 24;
    var msPerMonth = msPerDay * 30;
    var msPerYear = msPerDay * 365;

    var elapsed = Date.now() - timestamp;

    if (elapsed < msPerSecond) {
        return 'just now'
    } else if (elapsed < msPerMinute) {
        return Math.round(elapsed / 1000) + 's ago'
    } else if (elapsed < msPerHour) {
        return Math.round(elapsed / msPerMinute) + 'm ago'
    } else if (elapsed < msPerDay) {
        return Math.round(elapsed / msPerHour) + 'h ago'
    } else if (elapsed < msPerMonth) {
        return Math.round(elapsed / msPerDay) + 'd ago'
    } else if (elapsed < msPerYear) {
        return Math.round(elapsed / msPerMonth) + 'm ago'
    } else {
        return Math.round(elapsed / msPerYear) + 'y ago'
    }
}