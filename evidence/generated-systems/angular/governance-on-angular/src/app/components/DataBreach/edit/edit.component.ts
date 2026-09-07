import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataBreachService } from '../../../services/DataBreach.service';
import { SubBaseComponent } from '../../DataBreach/sub.base.component';


@Component({
    selector: 'app-edit-dataBreach',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataBreachComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataBreach';

    dataBreachForm: FormGroup;
    dataBreach: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataBreachService,
        private fb: FormBuilder
) {
        super(http);
        this.dataBreachForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  incidentDate: ['', Validators.required],
      description: ['', Validators.required],
      recordsAffected: ['', Validators.required],
      notificationRequired: ['', Validators.required],
      Organization: ['', ],
      ProcessingActivities: ['', ],
      DataCategories: ['', ],
      ThirdParties: ['', ],
      Matter: ['', ],
      Severity: ['', ],
      Status: ['', ]
        });
    }

    
    updateDataBreach(incidentDate, description, recordsAffected, notificationRequired, Organization, ProcessingActivities, DataCategories, ThirdParties, Matter, Severity, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataBreach(incidentDate, description, recordsAffected, notificationRequired, Organization, ProcessingActivities, DataCategories, ThirdParties, Matter, Severity, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataBreach']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataBreach(params['id']).subscribe(res => {
                this.dataBreach = res;
            });
        });
    }
}