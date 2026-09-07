import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RegulationService } from '../../../services/Regulation.service';
import { SubBaseComponent } from '../../Regulation/sub.base.component';


@Component({
    selector: 'app-edit-regulation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRegulationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Regulation';

    regulationForm: FormGroup;
    regulation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RegulationService,
        private fb: FormBuilder
) {
        super(http);
        this.regulationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      citation: ['', Validators.required],
      jurisdiction: ['', Validators.required],
      publicationUrl: ['', Validators.required],
      Obligations: ['', ],
      CompliancePrograms: ['', ]
        });
    }

    
    updateRegulation(name, citation, jurisdiction, publicationUrl, Obligations, CompliancePrograms): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRegulation(name, citation, jurisdiction, publicationUrl, Obligations, CompliancePrograms, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRegulation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRegulation(params['id']).subscribe(res => {
                this.regulation = res;
            });
        });
    }
}