import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DSPService } from '../../../services/DSP.service';
import { SubBaseComponent } from '../../DSP/sub.base.component';


@Component({
    selector: 'app-edit-dSP',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDSPComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DSP';

    dSPForm: FormGroup;
    dSP: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DSPService,
        private fb: FormBuilder
) {
        super(http);
        this.dSPForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      region: ['', Validators.required],
      AdAccounts: ['', ]
        });
    }

    
    updateDSP(name, website, region, AdAccounts): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDSP(name, website, region, AdAccounts, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDSP']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDSP(params['id']).subscribe(res => {
                this.dSP = res;
            });
        });
    }
}