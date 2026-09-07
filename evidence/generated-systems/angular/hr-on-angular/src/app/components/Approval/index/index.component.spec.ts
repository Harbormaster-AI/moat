
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexApprovalComponent } from './index.component';
import { ApprovalService } from '../../../services/Approval.service';

describe('IndexApprovalComponent', () => {
  let component: IndexApprovalComponent;
  let fixture: ComponentFixture<IndexApprovalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexApprovalComponent
      ],
      providers: [
        ApprovalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexApprovalComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});