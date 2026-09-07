import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataProviderService } from '../../../services/DataProvider.service';
import { SubBaseComponent } from '../../DataProvider/sub.base.component';


@Component({
    selector: 'app-edit-dataProvider',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataProviderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataProvider';

    dataProviderForm: FormGroup;
    dataProvider: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataProviderService,
        private fb: FormBuilder
) {
        super(http);
        this.dataProviderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      AudienceSegments: ['', ],
      ProviderType: ['', ]
        });
    }

    
    updateDataProvider(name, website, AudienceSegments, ProviderType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataProvider(name, website, AudienceSegments, ProviderType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataProvider']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataProvider(params['id']).subscribe(res => {
                this.dataProvider = res;
            });
        });
    }
}