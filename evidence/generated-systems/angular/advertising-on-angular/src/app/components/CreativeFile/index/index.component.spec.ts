
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCreativeFileComponent } from './index.component';
import { CreativeFileService } from '../../../services/CreativeFile.service';

describe('IndexCreativeFileComponent', () => {
  let component: IndexCreativeFileComponent;
  let fixture: ComponentFixture<IndexCreativeFileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCreativeFileComponent
      ],
      providers: [
        CreativeFileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCreativeFileComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});