import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataBreachService } from '../../../services/DataBreach.service';
import { DataBreach } from '../../../models/DataBreach';
import { SubBaseComponent } from '../../DataBreach/sub.base.component';

@Component({
    selector: 'app-create-dataBreach',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataBreachComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataBreach';

    dataBreachForm: FormGroup;
    dataBreach: DataBreach;

    constructor( http: HttpClient,
        private dataBreachService: DataBreachService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDataBreach(incidentDate, description, recordsAffected, notificationRequired, Organization, ProcessingActivities, DataCategories, ThirdParties, Matter, Severity, Status): void {
        this.dataBreachService
        .addDataBreach(incidentDate, description, recordsAffected, notificationRequired, Organization, ProcessingActivities, DataCategories, ThirdParties, Matter, Severity, Status)
            .subscribe(() => {
                this.router.navigate(['/indexDataBreach']);
            });
    }

    ngOnInit(): void {
    }
}