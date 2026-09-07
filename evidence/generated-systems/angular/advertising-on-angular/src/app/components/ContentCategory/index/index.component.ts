
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ContentCategoryService } from '../../../services/ContentCategory.service';
import { ContentCategory } from '../../../models/ContentCategory';

@Component({
    selector: 'app-index-contentCategory',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexContentCategoryComponent implements OnInit {

    contentCategorys: ContentCategory[] = [];

    constructor(
        private router: Router,
        private service: ContentCategoryService
) {}

    ngOnInit(): void {
        this.getContentCategorys();
}

    getContentCategorys(): void {
        this.service.getContentCategorys().subscribe((res) => {
        this.contentCategorys = res;
    });
}

    deleteContentCategory(id: any): void {
        this.service.deleteContentCategory(id)
            .subscribe(() => {
                this.getContentCategorys();
            });
    }
}