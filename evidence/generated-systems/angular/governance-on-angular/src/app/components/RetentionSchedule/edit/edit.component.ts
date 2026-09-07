import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RetentionScheduleService } from '../../../services/RetentionSchedule.service';
import { SubBaseComponent } from '../../RetentionSchedule/sub.base.component';


@Component({
    selector: 'app-edit-retentionSchedule',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRetentionScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RetentionSchedule';

    retentionScheduleForm: FormGroup;
    retentionSchedule: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RetentionScheduleService,
        private fb: FormBuilder
) {
        super(http);
        this.retentionScheduleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      retentionPeriodMonths: ['', Validators.required],
      Repositories: ['', ],
      Records: ['', ],
      Exceptions: ['', ],
      DispositionReviews: ['', ],
      RetentionTrigger: ['', ],
      DispositionAction: ['', ],
      Status: ['', ]
        });
    }

    
    updateRetentionSchedule(name, retentionPeriodMonths, Repositories, Records, Exceptions, DispositionReviews, RetentionTrigger, DispositionAction, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRetentionSchedule(name, retentionPeriodMonths, Repositories, Records, Exceptions, DispositionReviews, RetentionTrigger, DispositionAction, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRetentionSchedule']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRetentionSchedule(params['id']).subscribe(res => {
                this.retentionSchedule = res;
            });
        });
    }
}