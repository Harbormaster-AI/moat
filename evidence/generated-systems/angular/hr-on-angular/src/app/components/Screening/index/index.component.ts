
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ScreeningService } from '../../../services/Screening.service';
import { Screening } from '../../../models/Screening';

@Component({
    selector: 'app-index-screening',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexScreeningComponent implements OnInit {

    screenings: Screening[] = [];

    constructor(
        private router: Router,
        private service: ScreeningService
) {}

    ngOnInit(): void {
        this.getScreenings();
}

    getScreenings(): void {
        this.service.getScreenings().subscribe((res) => {
        this.screenings = res;
    });
}

    deleteScreening(id: any): void {
        this.service.deleteScreening(id)
            .subscribe(() => {
                this.getScreenings();
            });
    }
}