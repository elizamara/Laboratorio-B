document
  .getElementById("addItemForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();
    const itemValue = document.getElementById("itemValue").value;
    fetch("/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ value: itemValue }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Item added:", data);
        loadItems();
      });
  });

document
  .getElementById("removeItemForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();
    const itemId = parseInt(document.getElementById("itemId").value, 10);
    fetch("/remove", {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id: itemId }),
    }).then(() => {
      console.log("Item removed");
      loadItems();
    });
  });

function loadItems() {
  fetch("/list")
    .then((response) => response.json())
    .then((data) => {
      const itemsList = document.getElementById("itemsList");
      itemsList.innerHTML = "";
      data.forEach((item) => {
        const li = document.createElement("li");
        li.textContent = `ID: ${item.id}, Value: ${item.value}`;
        itemsList.appendChild(li);
      });
    });
}

loadItems();
