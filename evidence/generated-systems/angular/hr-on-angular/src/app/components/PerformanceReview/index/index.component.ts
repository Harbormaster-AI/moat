
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PerformanceReviewService } from '../../../services/PerformanceReview.service';
import { PerformanceReview } from '../../../models/PerformanceReview';

@Component({
    selector: 'app-index-performanceReview',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPerformanceReviewComponent implements OnInit {

    performanceReviews: PerformanceReview[] = [];

    constructor(
        private router: Router,
        private service: PerformanceReviewService
) {}

    ngOnInit(): void {
        this.getPerformanceReviews();
}

    getPerformanceReviews(): void {
        this.service.getPerformanceReviews().subscribe((res) => {
        this.performanceReviews = res;
    });
}

    deletePerformanceReview(id: any): void {
        this.service.deletePerformanceReview(id)
            .subscribe(() => {
                this.getPerformanceReviews();
            });
    }
}