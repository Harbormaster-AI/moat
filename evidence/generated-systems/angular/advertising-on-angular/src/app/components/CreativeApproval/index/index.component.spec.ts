
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCreativeApprovalComponent } from './index.component';
import { CreativeApprovalService } from '../../../services/CreativeApproval.service';

describe('IndexCreativeApprovalComponent', () => {
  let component: IndexCreativeApprovalComponent;
  let fixture: ComponentFixture<IndexCreativeApprovalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCreativeApprovalComponent
      ],
      providers: [
        CreativeApprovalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCreativeApprovalComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});