import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DispositionReviewService } from '../../../services/DispositionReview.service';
import { SubBaseComponent } from '../../DispositionReview/sub.base.component';


@Component({
    selector: 'app-edit-dispositionReview',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDispositionReviewComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DispositionReview';

    dispositionReviewForm: FormGroup;
    dispositionReview: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DispositionReviewService,
        private fb: FormBuilder
) {
        super(http);
        this.dispositionReviewForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reviewDate: ['', Validators.required],
      reviewer: ['', Validators.required],
      notes: ['', Validators.required],
      Record: ['', ],
      RetentionSchedule: ['', ],
      Outcome: ['', ]
        });
    }

    
    updateDispositionReview(reviewDate, reviewer, notes, Record, RetentionSchedule, Outcome): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDispositionReview(reviewDate, reviewer, notes, Record, RetentionSchedule, Outcome, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDispositionReview']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDispositionReview(params['id']).subscribe(res => {
                this.dispositionReview = res;
            });
        });
    }
}