
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditAuditWorkpaperComponent } from './edit.component';
import { AuditWorkpaperService } from '../../../services/AuditWorkpaper.service';

describe('EditAuditWorkpaperComponent', () => {
  let component: EditAuditWorkpaperComponent;
  let fixture: ComponentFixture<EditAuditWorkpaperComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditAuditWorkpaperComponent
      ],
      providers: [
        AuditWorkpaperService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditAuditWorkpaperComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});