import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataProcessingActivityService } from '../../../services/DataProcessingActivity.service';
import { DataProcessingActivity } from '../../../models/DataProcessingActivity';
import { SubBaseComponent } from '../../DataProcessingActivity/sub.base.component';

@Component({
    selector: 'app-create-dataProcessingActivity',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataProcessingActivityComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataProcessingActivity';

    dataProcessingActivityForm: FormGroup;
    dataProcessingActivity: DataProcessingActivity;

    constructor( http: HttpClient,
        private dataProcessingActivityService: DataProcessingActivityService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dataProcessingActivityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      purpose: ['', Validators.required],
      startDate: ['', Validators.required],
      Organization: ['', ],
      DataCategories: ['', ],
      Systems: ['', ],
      Records: ['', ],
      PrivacyNotices: ['', ],
      ThirdParties: ['', ],
      Consents: ['', ],
      DataBreaches: ['', ],
      DataSubjectRequests: ['', ],
      LawfulBasis: ['', ]
        });
    }

    
    addDataProcessingActivity(name, purpose, startDate, Organization, DataCategories, Systems, Records, PrivacyNotices, ThirdParties, Consents, DataBreaches, DataSubjectRequests, LawfulBasis): void {
        this.dataProcessingActivityService
        .addDataProcessingActivity(name, purpose, startDate, Organization, DataCategories, Systems, Records, PrivacyNotices, ThirdParties, Consents, DataBreaches, DataSubjectRequests, LawfulBasis)
            .subscribe(() => {
                this.router.navigate(['/indexDataProcessingActivity']);
            });
    }

    ngOnInit(): void {
    }
}