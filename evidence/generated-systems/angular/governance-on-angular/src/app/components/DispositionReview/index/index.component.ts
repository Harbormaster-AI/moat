
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DispositionReviewService } from '../../../services/DispositionReview.service';
import { DispositionReview } from '../../../models/DispositionReview';

@Component({
    selector: 'app-index-dispositionReview',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDispositionReviewComponent implements OnInit {

    dispositionReviews: DispositionReview[] = [];

    constructor(
        private router: Router,
        private service: DispositionReviewService
) {}

    ngOnInit(): void {
        this.getDispositionReviews();
}

    getDispositionReviews(): void {
        this.service.getDispositionReviews().subscribe((res) => {
        this.dispositionReviews = res;
    });
}

    deleteDispositionReview(id: any): void {
        this.service.deleteDispositionReview(id)
            .subscribe(() => {
                this.getDispositionReviews();
            });
    }
}