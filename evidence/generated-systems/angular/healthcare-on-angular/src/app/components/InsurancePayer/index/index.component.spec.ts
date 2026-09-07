
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInsurancePayerComponent } from './index.component';
import { InsurancePayerService } from '../../../services/InsurancePayer.service';

describe('IndexInsurancePayerComponent', () => {
  let component: IndexInsurancePayerComponent;
  let fixture: ComponentFixture<IndexInsurancePayerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInsurancePayerComponent
      ],
      providers: [
        InsurancePayerService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInsurancePayerComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});