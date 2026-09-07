
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexReinsuranceAgreementComponent } from './index.component';
import { ReinsuranceAgreementService } from '../../../services/ReinsuranceAgreement.service';

describe('IndexReinsuranceAgreementComponent', () => {
  let component: IndexReinsuranceAgreementComponent;
  let fixture: ComponentFixture<IndexReinsuranceAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexReinsuranceAgreementComponent
      ],
      providers: [
        ReinsuranceAgreementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexReinsuranceAgreementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});