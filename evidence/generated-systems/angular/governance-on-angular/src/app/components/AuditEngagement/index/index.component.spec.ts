
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAuditEngagementComponent } from './index.component';
import { AuditEngagementService } from '../../../services/AuditEngagement.service';

describe('IndexAuditEngagementComponent', () => {
  let component: IndexAuditEngagementComponent;
  let fixture: ComponentFixture<IndexAuditEngagementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAuditEngagementComponent
      ],
      providers: [
        AuditEngagementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAuditEngagementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});