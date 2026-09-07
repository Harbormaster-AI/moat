import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivityService } from '../../../services/Activity.service';
import { Activity } from '../../../models/Activity';
import { SubBaseComponent } from '../../Activity/sub.base.component';

@Component({
    selector: 'app-create-activity',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateActivityComponent extends SubBaseComponent implements OnInit {

    title = 'Add Activity';

    activityForm: FormGroup;
    activity: Activity;

    constructor( http: HttpClient,
        private activityService: ActivityService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.activityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  subject: ['', Validators.required],
      dueDate: ['', Validators.required],
      startAt: ['', Validators.required],
      endAt: ['', Validators.required],
      location: ['', Validators.required],
      Organization: ['', ],
      Owner: ['', ],
      Account: ['', ],
      Contact: ['', ],
      Lead: ['', ],
      Opportunity: ['', ],
      Case: ['', ],
      Campaign: ['', ],
      ActivityType: ['', ],
      Status: ['', ],
      Priority: ['', ]
        });
    }

    
    addActivity(subject, dueDate, startAt, endAt, location, Organization, Owner, Account, Contact, Lead, Opportunity, Case, Campaign, ActivityType, Status, Priority): void {
        this.activityService
        .addActivity(subject, dueDate, startAt, endAt, location, Organization, Owner, Account, Contact, Lead, Opportunity, Case, Campaign, ActivityType, Status, Priority)
            .subscribe(() => {
                this.router.navigate(['/indexActivity']);
            });
    }

    ngOnInit(): void {
    }
}