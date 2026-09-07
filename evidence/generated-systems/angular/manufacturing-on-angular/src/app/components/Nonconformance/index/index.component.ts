
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NonconformanceService } from '../../../services/Nonconformance.service';
import { Nonconformance } from '../../../models/Nonconformance';

@Component({
    selector: 'app-index-nonconformance',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexNonconformanceComponent implements OnInit {

    nonconformances: Nonconformance[] = [];

    constructor(
        private router: Router,
        private service: NonconformanceService
) {}

    ngOnInit(): void {
        this.getNonconformances();
}

    getNonconformances(): void {
        this.service.getNonconformances().subscribe((res) => {
        this.nonconformances = res;
    });
}

    deleteNonconformance(id: any): void {
        this.service.deleteNonconformance(id)
            .subscribe(() => {
                this.getNonconformances();
            });
    }
}