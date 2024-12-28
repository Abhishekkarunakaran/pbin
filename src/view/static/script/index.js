
const copyContent = async () => {
    let text = document.getElementById('resultTab').innerHTML;
    let button = document.getElementById('copyButton');
    try {
        await navigator.clipboard.writeText(text);
        button.innerHTML = "Copied!";
        console.log('Content copied to clipboard')
    } catch (err) {
        console.error('Failed to copy: ', err);
    }
}