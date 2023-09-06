const API_URL = 'http://localhost:3000';

async function uploadFile(topics, fileElement){
    const formData = new FormData();

    formData.append("file", fileElement.files[0]);
    formData.append("topics", topics);

    await fetch(`${API_URL}/upload`, {
        method: 'post',
        body: formData,
    });
}

async function sendQuery(topics, dateStart, dateEnd) {


    const params = encodeURIComponent(topics.join())
    const dateStartE = encodeURIComponent(dateStart)
    const dateEndE = encodeURIComponent(dateEnd)

    const r = await fetch(`${API_URL}/search?query=${params}&dateStart=${dateStartE}&dateEnd=${dateEndE}`);
    const data = await r.json()
    console.log(data);
    return sortObjKeysByValue(data.Data)
}

//extremely naive and definetly slow way to accomplish sorting object by value adn return keys as array
function sortObjKeysByValue(obj){
    const results = []

    let max_vals = [...new Set(Object.values(obj).sort((a, b) => (a-b)).reverse())];
    for(let i = 0; i < max_vals.length; i++){
        let cur_max = max_vals[i]
        for(const [key, value] of Object.entries(obj)){
            if(value == cur_max){
                results.push(key)
            }
        }
    }

    return results
}