import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ActivityService } from '../../../services/Activity.service';
import { SubBaseComponent } from '../../Activity/sub.base.component';


@Component({
    selector: 'app-edit-activity',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditActivityComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Activity';

    activityForm: FormGroup;
    activity: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ActivityService,
        private fb: FormBuilder
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

    
    updateActivity(subject, dueDate, startAt, endAt, location, Organization, Owner, Account, Contact, Lead, Opportunity, Case, Campaign, ActivityType, Status, Priority): void {
        this.route.params.subscribe((params) => {

                        this.service.updateActivity(subject, dueDate, startAt, endAt, location, Organization, Owner, Account, Contact, Lead, Opportunity, Case, Campaign, ActivityType, Status, Priority, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexActivity']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getActivity(params['id']).subscribe(res => {
                this.activity = res;
            });
        });
    }
}