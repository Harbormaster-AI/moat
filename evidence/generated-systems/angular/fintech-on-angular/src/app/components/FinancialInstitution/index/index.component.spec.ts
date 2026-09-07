
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFinancialInstitutionComponent } from './index.component';
import { FinancialInstitutionService } from '../../../services/FinancialInstitution.service';

describe('IndexFinancialInstitutionComponent', () => {
  let component: IndexFinancialInstitutionComponent;
  let fixture: ComponentFixture<IndexFinancialInstitutionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFinancialInstitutionComponent
      ],
      providers: [
        FinancialInstitutionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFinancialInstitutionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});