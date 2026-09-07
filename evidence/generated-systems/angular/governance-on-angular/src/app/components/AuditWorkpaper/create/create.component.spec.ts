
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAuditWorkpaperComponent } from './create.component';
import { AuditWorkpaperService } from '../../../services/AuditWorkpaper.service';
import { Router } from '@angular/router';

describe('CreateAuditWorkpaperComponent', () => {
  let component: CreateAuditWorkpaperComponent;
  let fixture: ComponentFixture<CreateAuditWorkpaperComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAuditWorkpaperComponent
      ],
      providers: [
        AuditWorkpaperService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAuditWorkpaperComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});