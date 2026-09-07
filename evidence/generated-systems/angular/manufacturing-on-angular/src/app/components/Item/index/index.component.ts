
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ItemService } from '../../../services/Item.service';
import { Item } from '../../../models/Item';

@Component({
    selector: 'app-index-item',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexItemComponent implements OnInit {

    items: Item[] = [];

    constructor(
        private router: Router,
        private service: ItemService
) {}

    ngOnInit(): void {
        this.getItems();
}

    getItems(): void {
        this.service.getItems().subscribe((res) => {
        this.items = res;
    });
}

    deleteItem(id: any): void {
        this.service.deleteItem(id)
            .subscribe(() => {
                this.getItems();
            });
    }
}