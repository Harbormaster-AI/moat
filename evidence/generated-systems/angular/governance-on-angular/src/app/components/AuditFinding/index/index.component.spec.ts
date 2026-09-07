
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAuditFindingComponent } from './index.component';
import { AuditFindingService } from '../../../services/AuditFinding.service';

describe('IndexAuditFindingComponent', () => {
  let component: IndexAuditFindingComponent;
  let fixture: ComponentFixture<IndexAuditFindingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAuditFindingComponent
      ],
      providers: [
        AuditFindingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAuditFindingComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});