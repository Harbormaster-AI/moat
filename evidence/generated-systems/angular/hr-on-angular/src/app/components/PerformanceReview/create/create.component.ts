import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PerformanceReviewService } from '../../../services/PerformanceReview.service';
import { PerformanceReview } from '../../../models/PerformanceReview';
import { SubBaseComponent } from '../../PerformanceReview/sub.base.component';

@Component({
    selector: 'app-create-performanceReview',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePerformanceReviewComponent extends SubBaseComponent implements OnInit {

    title = 'Add PerformanceReview';

    performanceReviewForm: FormGroup;
    performanceReview: PerformanceReview;

    constructor( http: HttpClient,
        private performanceReviewService: PerformanceReviewService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.performanceReviewForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reviewNumber: ['', Validators.required],
      reviewDate: ['', Validators.required],
      reviewerComments: ['', Validators.required],
      Employee: ['', ],
      Reviewer: ['', ],
      Cycle: ['', ],
      CompetencyRatings: ['', ],
      Goals: ['', ],
      Rating: ['', ],
      Status: ['', ]
        });
    }

    
    addPerformanceReview(reviewNumber, reviewDate, reviewerComments, Employee, Reviewer, Cycle, CompetencyRatings, Goals, Rating, Status): void {
        this.performanceReviewService
        .addPerformanceReview(reviewNumber, reviewDate, reviewerComments, Employee, Reviewer, Cycle, CompetencyRatings, Goals, Rating, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPerformanceReview']);
            });
    }

    ngOnInit(): void {
    }
}