import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PerformanceReviewService } from '../../../services/PerformanceReview.service';
import { SubBaseComponent } from '../../PerformanceReview/sub.base.component';


@Component({
    selector: 'app-edit-performanceReview',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPerformanceReviewComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PerformanceReview';

    performanceReviewForm: FormGroup;
    performanceReview: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PerformanceReviewService,
        private fb: FormBuilder
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

    
    updatePerformanceReview(reviewNumber, reviewDate, reviewerComments, Employee, Reviewer, Cycle, CompetencyRatings, Goals, Rating, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePerformanceReview(reviewNumber, reviewDate, reviewerComments, Employee, Reviewer, Cycle, CompetencyRatings, Goals, Rating, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPerformanceReview']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPerformanceReview(params['id']).subscribe(res => {
                this.performanceReview = res;
            });
        });
    }
}