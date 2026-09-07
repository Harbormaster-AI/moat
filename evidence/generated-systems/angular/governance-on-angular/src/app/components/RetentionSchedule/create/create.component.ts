import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RetentionScheduleService } from '../../../services/RetentionSchedule.service';
import { RetentionSchedule } from '../../../models/RetentionSchedule';
import { SubBaseComponent } from '../../RetentionSchedule/sub.base.component';

@Component({
    selector: 'app-create-retentionSchedule',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRetentionScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Add RetentionSchedule';

    retentionScheduleForm: FormGroup;
    retentionSchedule: RetentionSchedule;

    constructor( http: HttpClient,
        private retentionScheduleService: RetentionScheduleService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRetentionSchedule(name, retentionPeriodMonths, Repositories, Records, Exceptions, DispositionReviews, RetentionTrigger, DispositionAction, Status): void {
        this.retentionScheduleService
        .addRetentionSchedule(name, retentionPeriodMonths, Repositories, Records, Exceptions, DispositionReviews, RetentionTrigger, DispositionAction, Status)
            .subscribe(() => {
                this.router.navigate(['/indexRetentionSchedule']);
            });
    }

    ngOnInit(): void {
    }
}