
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAuditProgramComponent } from './index.component';
import { AuditProgramService } from '../../../services/AuditProgram.service';

describe('IndexAuditProgramComponent', () => {
  let component: IndexAuditProgramComponent;
  let fixture: ComponentFixture<IndexAuditProgramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAuditProgramComponent
      ],
      providers: [
        AuditProgramService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAuditProgramComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});