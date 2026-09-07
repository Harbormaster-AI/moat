import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OnboardingTaskService } from '../../../services/OnboardingTask.service';
import { OnboardingTask } from '../../../models/OnboardingTask';
import { SubBaseComponent } from '../../OnboardingTask/sub.base.component';

@Component({
    selector: 'app-create-onboardingTask',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOnboardingTaskComponent extends SubBaseComponent implements OnInit {

    title = 'Add OnboardingTask';

    onboardingTaskForm: FormGroup;
    onboardingTask: OnboardingTask;

    constructor( http: HttpClient,
        private onboardingTaskService: OnboardingTaskService,
        private fb: FormBuilder,
        private router: Router
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

    
    addOnboardingTask(taskNumber, name, dueDate, Employee, AssignedTo, Dependencies, RelatedOffer, Status): void {
        this.onboardingTaskService
        .addOnboardingTask(taskNumber, name, dueDate, Employee, AssignedTo, Dependencies, RelatedOffer, Status)
            .subscribe(() => {
                this.router.navigate(['/indexOnboardingTask']);
            });
    }

    ngOnInit(): void {
    }
}