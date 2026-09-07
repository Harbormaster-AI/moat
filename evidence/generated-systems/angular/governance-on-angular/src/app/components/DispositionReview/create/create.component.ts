import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DispositionReviewService } from '../../../services/DispositionReview.service';
import { DispositionReview } from '../../../models/DispositionReview';
import { SubBaseComponent } from '../../DispositionReview/sub.base.component';

@Component({
    selector: 'app-create-dispositionReview',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDispositionReviewComponent extends SubBaseComponent implements OnInit {

    title = 'Add DispositionReview';

    dispositionReviewForm: FormGroup;
    dispositionReview: DispositionReview;

    constructor( http: HttpClient,
        private dispositionReviewService: DispositionReviewService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDispositionReview(reviewDate, reviewer, notes, Record, RetentionSchedule, Outcome): void {
        this.dispositionReviewService
        .addDispositionReview(reviewDate, reviewer, notes, Record, RetentionSchedule, Outcome)
            .subscribe(() => {
                this.router.navigate(['/indexDispositionReview']);
            });
    }

    ngOnInit(): void {
    }
}