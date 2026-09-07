
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEmploymentContractComponent } from './index.component';
import { EmploymentContractService } from '../../../services/EmploymentContract.service';

describe('IndexEmploymentContractComponent', () => {
  let component: IndexEmploymentContractComponent;
  let fixture: ComponentFixture<IndexEmploymentContractComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEmploymentContractComponent
      ],
      providers: [
        EmploymentContractService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEmploymentContractComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});