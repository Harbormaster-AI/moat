import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SoftwareLoadService } from '../../../services/SoftwareLoad.service';
import { SubBaseComponent } from '../../SoftwareLoad/sub.base.component';


@Component({
    selector: 'app-edit-softwareLoad',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSoftwareLoadComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SoftwareLoad';

    softwareLoadForm: FormGroup;
    softwareLoad: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SoftwareLoadService,
        private fb: FormBuilder
) {
        super(http);
        this.softwareLoadForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  version: ['', Validators.required],
      ConnectedAircraft: ['', ],
      AvionicsSuite: ['', ],
      LoadType: ['', ]
        });
    }

    
    updateSoftwareLoad(version, ConnectedAircraft, AvionicsSuite, LoadType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSoftwareLoad(version, ConnectedAircraft, AvionicsSuite, LoadType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSoftwareLoad']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSoftwareLoad(params['id']).subscribe(res => {
                this.softwareLoad = res;
            });
        });
    }
}