import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RegulationService } from '../../../services/Regulation.service';
import { Regulation } from '../../../models/Regulation';
import { SubBaseComponent } from '../../Regulation/sub.base.component';

@Component({
    selector: 'app-create-regulation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRegulationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Regulation';

    regulationForm: FormGroup;
    regulation: Regulation;

    constructor( http: HttpClient,
        private regulationService: RegulationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRegulation(name, citation, jurisdiction, publicationUrl, Obligations, CompliancePrograms): void {
        this.regulationService
        .addRegulation(name, citation, jurisdiction, publicationUrl, Obligations, CompliancePrograms)
            .subscribe(() => {
                this.router.navigate(['/indexRegulation']);
            });
    }

    ngOnInit(): void {
    }
}