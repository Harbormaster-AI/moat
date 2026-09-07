
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataCategoryService } from '../../../services/DataCategory.service';
import { DataCategory } from '../../../models/DataCategory';

@Component({
    selector: 'app-index-dataCategory',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataCategoryComponent implements OnInit {

    dataCategorys: DataCategory[] = [];

    constructor(
        private router: Router,
        private service: DataCategoryService
) {}

    ngOnInit(): void {
        this.getDataCategorys();
}

    getDataCategorys(): void {
        this.service.getDataCategorys().subscribe((res) => {
        this.dataCategorys = res;
    });
}

    deleteDataCategory(id: any): void {
        this.service.deleteDataCategory(id)
            .subscribe(() => {
                this.getDataCategorys();
            });
    }
}