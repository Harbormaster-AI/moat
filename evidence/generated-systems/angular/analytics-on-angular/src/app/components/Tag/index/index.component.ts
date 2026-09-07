
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TagService } from '../../../services/Tag.service';
import { Tag } from '../../../models/Tag';

@Component({
    selector: 'app-index-tag',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTagComponent implements OnInit {

    tags: Tag[] = [];

    constructor(
        private router: Router,
        private service: TagService
) {}

    ngOnInit(): void {
        this.getTags();
}

    getTags(): void {
        this.service.getTags().subscribe((res) => {
        this.tags = res;
    });
}

    deleteTag(id: any): void {
        this.service.deleteTag(id)
            .subscribe(() => {
                this.getTags();
            });
    }
}