
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexModelVersionComponent } from './index.component';
import { ModelVersionService } from '../../../services/ModelVersion.service';

describe('IndexModelVersionComponent', () => {
  let component: IndexModelVersionComponent;
  let fixture: ComponentFixture<IndexModelVersionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexModelVersionComponent
      ],
      providers: [
        ModelVersionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexModelVersionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});