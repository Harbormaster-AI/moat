import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OnboardingTaskService } from '../../../services/OnboardingTask.service';
import { SubBaseComponent } from '../../OnboardingTask/sub.base.component';


@Component({
    selector: 'app-edit-onboardingTask',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOnboardingTaskComponent extends SubBaseComponent implements OnInit {

    title = 'Edit OnboardingTask';

    onboardingTaskForm: FormGroup;
    onboardingTask: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OnboardingTaskService,
        private fb: FormBuilder
) {
        super(http);
        this.onboardingTaskForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  taskNumber: ['', Validators.required],
      name: ['', Validators.required],
      dueDate: ['', Validators.required],
      Employee: ['', ],
      AssignedTo: ['', ],
      Dependencies: ['', ],
      RelatedOffer: ['', ],
      Status: ['', ]
        });
    }

    
    updateOnboardingTask(taskNumber, name, dueDate, Employee, AssignedTo, Dependencies, RelatedOffer, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOnboardingTask(taskNumber, name, dueDate, Employee, AssignedTo, Dependencies, RelatedOffer, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOnboardingTask']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOnboardingTask(params['id']).subscribe(res => {
                this.onboardingTask = res;
            });
        });
    }
}