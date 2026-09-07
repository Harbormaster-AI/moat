
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAuditWorkpaperComponent } from './index.component';
import { AuditWorkpaperService } from '../../../services/AuditWorkpaper.service';

describe('IndexAuditWorkpaperComponent', () => {
  let component: IndexAuditWorkpaperComponent;
  let fixture: ComponentFixture<IndexAuditWorkpaperComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAuditWorkpaperComponent
      ],
      providers: [
        AuditWorkpaperService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAuditWorkpaperComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});