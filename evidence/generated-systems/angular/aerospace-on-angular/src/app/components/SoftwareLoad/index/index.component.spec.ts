
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSoftwareLoadComponent } from './index.component';
import { SoftwareLoadService } from '../../../services/SoftwareLoad.service';

describe('IndexSoftwareLoadComponent', () => {
  let component: IndexSoftwareLoadComponent;
  let fixture: ComponentFixture<IndexSoftwareLoadComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSoftwareLoadComponent
      ],
      providers: [
        SoftwareLoadService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSoftwareLoadComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});