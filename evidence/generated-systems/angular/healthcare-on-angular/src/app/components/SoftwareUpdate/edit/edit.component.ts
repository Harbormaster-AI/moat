import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SoftwareUpdateService } from '../../../services/SoftwareUpdate.service';
import { SubBaseComponent } from '../../SoftwareUpdate/sub.base.component';


@Component({
    selector: 'app-edit-softwareUpdate',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSoftwareUpdateComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SoftwareUpdate';

    softwareUpdateForm: FormGroup;
    softwareUpdate: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SoftwareUpdateService,
        private fb: FormBuilder
) {
        super(http);
        this.softwareUpdateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  version: ['', Validators.required],
      appliedDate: ['', Validators.required],
      Device: ['', ],
      UpdateType: ['', ]
        });
    }

    
    updateSoftwareUpdate(version, appliedDate, Device, UpdateType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSoftwareUpdate(version, appliedDate, Device, UpdateType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSoftwareUpdate']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSoftwareUpdate(params['id']).subscribe(res => {
                this.softwareUpdate = res;
            });
        });
    }
}