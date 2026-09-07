
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AdjusterService } from '../../../services/Adjuster.service';
import { Adjuster } from '../../../models/Adjuster';

@Component({
    selector: 'app-index-adjuster',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAdjusterComponent implements OnInit {

    adjusters: Adjuster[] = [];

    constructor(
        private router: Router,
        private service: AdjusterService
) {}

    ngOnInit(): void {
        this.getAdjusters();
}

    getAdjusters(): void {
        this.service.getAdjusters().subscribe((res) => {
        this.adjusters = res;
    });
}

    deleteAdjuster(id: any): void {
        this.service.deleteAdjuster(id)
            .subscribe(() => {
                this.getAdjusters();
            });
    }
}