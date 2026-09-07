import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OrganizationService } from '../../../services/Organization.service';
import { Organization } from '../../../models/Organization';
import { SubBaseComponent } from '../../Organization/sub.base.component';

@Component({
    selector: 'app-create-organization',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOrganizationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Organization';

    organizationForm: FormGroup;
    organization: Organization;

    constructor( http: HttpClient,
        private organizationService: OrganizationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.organizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      jurisdiction: ['', Validators.required],
      industrySector: ['', Validators.required],
      GovernanceBodies: ['', ],
      Policies: ['', ],
      Risks: ['', ],
      ThirdParties: ['', ],
      RecordsRepositories: ['', ],
      DataProcessingActivities: ['', ],
      CompliancePrograms: ['', ],
      AuditPrograms: ['', ],
      BusinessUnits: ['', ],
      Matters: ['', ],
      DataBreaches: ['', ]
        });
    }

    
    addOrganization(name, legalName, jurisdiction, industrySector, GovernanceBodies, Policies, Risks, ThirdParties, RecordsRepositories, DataProcessingActivities, CompliancePrograms, AuditPrograms, BusinessUnits, Matters, DataBreaches): void {
        this.organizationService
        .addOrganization(name, legalName, jurisdiction, industrySector, GovernanceBodies, Policies, Risks, ThirdParties, RecordsRepositories, DataProcessingActivities, CompliancePrograms, AuditPrograms, BusinessUnits, Matters, DataBreaches)
            .subscribe(() => {
                this.router.navigate(['/indexOrganization']);
            });
    }

    ngOnInit(): void {
    }
}