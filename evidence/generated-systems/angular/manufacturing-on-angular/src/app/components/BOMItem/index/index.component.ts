
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BOMItemService } from '../../../services/BOMItem.service';
import { BOMItem } from '../../../models/BOMItem';

@Component({
    selector: 'app-index-bOMItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBOMItemComponent implements OnInit {

    bOMItems: BOMItem[] = [];

    constructor(
        private router: Router,
        private service: BOMItemService
) {}

    ngOnInit(): void {
        this.getBOMItems();
}

    getBOMItems(): void {
        this.service.getBOMItems().subscribe((res) => {
        this.bOMItems = res;
    });
}

    deleteBOMItem(id: any): void {
        this.service.deleteBOMItem(id)
            .subscribe(() => {
                this.getBOMItems();
            });
    }
}