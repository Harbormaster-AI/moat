
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexComplianceAlertComponent } from './index.component';
import { ComplianceAlertService } from '../../../services/ComplianceAlert.service';

describe('IndexComplianceAlertComponent', () => {
  let component: IndexComplianceAlertComponent;
  let fixture: ComponentFixture<IndexComplianceAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexComplianceAlertComponent
      ],
      providers: [
        ComplianceAlertService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexComplianceAlertComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});