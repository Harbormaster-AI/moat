
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LineItemService } from '../../../services/LineItem.service';
import { LineItem } from '../../../models/LineItem';

@Component({
    selector: 'app-index-lineItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLineItemComponent implements OnInit {

    lineItems: LineItem[] = [];

    constructor(
        private router: Router,
        private service: LineItemService
) {}

    ngOnInit(): void {
        this.getLineItems();
}

    getLineItems(): void {
        this.service.getLineItems().subscribe((res) => {
        this.lineItems = res;
    });
}

    deleteLineItem(id: any): void {
        this.service.deleteLineItem(id)
            .subscribe(() => {
                this.getLineItems();
            });
    }
}