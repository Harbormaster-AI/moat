
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAppliedFeeComponent } from './index.component';
import { AppliedFeeService } from '../../../services/AppliedFee.service';

describe('IndexAppliedFeeComponent', () => {
  let component: IndexAppliedFeeComponent;
  let fixture: ComponentFixture<IndexAppliedFeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAppliedFeeComponent
      ],
      providers: [
        AppliedFeeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAppliedFeeComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});