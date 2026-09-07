import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AnalyticsWorkspaceService } from '../../../services/AnalyticsWorkspace.service';
import { AnalyticsWorkspace } from '../../../models/AnalyticsWorkspace';
import { SubBaseComponent } from '../../AnalyticsWorkspace/sub.base.component';

@Component({
    selector: 'app-create-analyticsWorkspace',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAnalyticsWorkspaceComponent extends SubBaseComponent implements OnInit {

    title = 'Add AnalyticsWorkspace';

    analyticsWorkspaceForm: FormGroup;
    analyticsWorkspace: AnalyticsWorkspace;

    constructor( http: HttpClient,
        private analyticsWorkspaceService: AnalyticsWorkspaceService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.analyticsWorkspaceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      businessDomain: ['', Validators.required],
      ownerTeam: ['', Validators.required],
      Datasets: ['', ],
      DataSources: ['', ],
      Pipelines: ['', ],
      Dashboards: ['', ],
      Reports: ['', ],
      Notebooks: ['', ],
      Models: ['', ],
      FeatureSets: ['', ],
      Policies: ['', ],
      LineageNodes: ['', ],
      GovernanceTier: ['', ]
        });
    }

    
    addAnalyticsWorkspace(name, businessDomain, ownerTeam, Datasets, DataSources, Pipelines, Dashboards, Reports, Notebooks, Models, FeatureSets, Policies, LineageNodes, GovernanceTier): void {
        this.analyticsWorkspaceService
        .addAnalyticsWorkspace(name, businessDomain, ownerTeam, Datasets, DataSources, Pipelines, Dashboards, Reports, Notebooks, Models, FeatureSets, Policies, LineageNodes, GovernanceTier)
            .subscribe(() => {
                this.router.navigate(['/indexAnalyticsWorkspace']);
            });
    }

    ngOnInit(): void {
    }
}