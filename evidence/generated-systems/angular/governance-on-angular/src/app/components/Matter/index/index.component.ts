
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MatterService } from '../../../services/Matter.service';
import { Matter } from '../../../models/Matter';

@Component({
    selector: 'app-index-matter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMatterComponent implements OnInit {

    matters: Matter[] = [];

    constructor(
        private router: Router,
        private service: MatterService
) {}

    ngOnInit(): void {
        this.getMatters();
}

    getMatters(): void {
        this.service.getMatters().subscribe((res) => {
        this.matters = res;
    });
}

    deleteMatter(id: any): void {
        this.service.deleteMatter(id)
            .subscribe(() => {
                this.getMatters();
            });
    }
}