import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DSPService } from '../../../services/DSP.service';
import { DSP } from '../../../models/DSP';
import { SubBaseComponent } from '../../DSP/sub.base.component';

@Component({
    selector: 'app-create-dSP',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDSPComponent extends SubBaseComponent implements OnInit {

    title = 'Add DSP';

    dSPForm: FormGroup;
    dSP: DSP;

    constructor( http: HttpClient,
        private dSPService: DSPService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDSP(name, website, region, AdAccounts): void {
        this.dSPService
        .addDSP(name, website, region, AdAccounts)
            .subscribe(() => {
                this.router.navigate(['/indexDSP']);
            });
    }

    ngOnInit(): void {
    }
}