function get(url){
    const Http = new XMLHttpRequest();
    Http.open("GET", url, false);
    Http.send(null);
    console.log(Http.status)
    var data = JSON.parse(Http.responseText) // var = let
    console.log(data)
    return data
}

function getTitle() {
    return get('https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/package.json').name  
}

function getDescription() {
    return get('https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/package.json').description
}

function getPage(){
    // Title
    title_rez = getTitle();
    // Description
    description_rez = getDescription();
    // hello_pic path
    
    // carousel_pics paths
    

}